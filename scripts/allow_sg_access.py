#!/usr/bin/env python3
import argparse
import boto3
from botocore.exceptions import ClientError


def parse_arguments():
    parser = argparse.ArgumentParser(description='grant access to certain security groups for vuln scanning')
    parser.add_argument('--region',default='us-east-1', help='Region to apply roles to')
    parser.add_argument('--vpc-id',help='VPC id to apply')
    parser.add_argument('--source-sg-id', help='Destination security group to open')
    parser.add_argument('--task', required=True, default="remove", help="Task is to grant or remove permissions")
    return parser.parse_args()


def get_all_sg_for_vpc(ec2, vpc_id):
    kwargs = {
        'Filters': [
            {
                'Name': 'vpc-id',
                'Values': [ vpc_id ]
            }
        ]}
    sgs = []
    done = False
    while not done:
        response = ec2.describe_security_groups(**kwargs)
        for sg in response['SecurityGroups']:
            sgs.append(sg)
        if 'NextToken' not in response:
            done = True
        else:
            kwargs['NextToken'] = response['NextToken']
    return sgs

def check_rule_exists(sg, account_id, group_id, from_port=0,to_port=65535,ip_protocol='-1'):
    rule_exists = False
    for perms in sg['IpPermissions']:
        if perms['IpProtocol'] != ip_protocol:
            continue
        if ip_protocol != '-1':
            if perms['FromPort'] != from_port and perm['ToPort'] != to_port:
                continue
        if rule_exists:
            break
        for group_pair in perms['UserIdGroupPairs']:
            if group_pair['UserId'] == account_id and group_pair['GroupId'] == group_id:
                rule_exists = True
                break
    return rule_exists

# Pick all security groups in the VPC and grant access to be scanned
def manage_access_to_sg(region, vpc_id, src_sg_id, task):
    ec2 = boto3.client('ec2', region_name=region)
    sg_ids = get_all_sg_for_vpc(ec2, vpc_id)
    if task == "grant":
        for sg in sg_ids:
            rule_exists = check_rule_exists(sg, "660610034966", src_sg_id)
            if not rule_exists:
                data = ec2.authorize_security_group_ingress(
                    GroupId=sg['GroupId'],

                    IpPermissions=[
                        {

                            'FromPort': 0,
                            'ToPort': 65535,
                            'IpProtocol': '-1',
                            'UserIdGroupPairs':[{
                                'Description': 'Grant access for pen-test',
                                'GroupId': src_sg_id
                            }]
                        }
                    ],
                    DryRun=False
                )
            print('Ingress rule successfully set for %s'% sg['GroupName'])
    elif task == "revoke":
        for sg in sg_ids:
            rule_exists = check_rule_exists(sg, "660610034966", src_sg_id)
            if rule_exists:
                print("rule exists")
                data = ec2.revoke_security_group_ingress(
                    GroupId=sg['GroupId'],
                    IpPermissions=[
                        {
                            'FromPort': 0,
                            'ToPort': 65535,
                            'IpProtocol': '-1',
                            'UserIdGroupPairs':[{
                                'GroupId': src_sg_id
                            }]
                        }
                    ],
                    DryRun=False
                )
                print('Ingress rule sucessfully revoked for %s'%sg['GroupName'])


def main():
    args = parse_arguments()
    manage_access_to_sg(args.region, args.vpc_id, args.source_sg_id, args.task)

if __name__ == '__main__':
    main()
