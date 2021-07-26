import boto3

# See https://boto3.amazonaws.com/v1/documentation/api/latest/reference/core/session.html
# See https://stackoverflow.com/questions/33378422/how-to-choose-an-aws-profile-when-using-boto3-to-connect-to-cloudfront
ses = boto3.session.Session(region_name='us-east-2')  # profile_name='your-profile'

client = ses.client('ec2')

# See https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/ec2.html#EC2.Client.run_instances
result = client.run_instances(
    ImageId='ami-0b29b6e62f2343b46',  # ImageId: https://cloud-images.ubuntu.com/locator/ec2/
    InstanceType='t3.small',
    # Optional parameters you likely want to define
    # KeyName: 'some-key',
    # SecurityGroupIds: = ['sg-foobar',]
    # SubnetId: 'subnet-foobar',

    MaxCount=1,
    MinCount=1,
    BlockDeviceMappings=[
        {
            'DeviceName': '/dev/sda1',
            'Ebs': {
                'DeleteOnTermination': True,
                'VolumeType': 'gp3',
                'VolumeSize': 8,
            }
        }
    ]
)

# See https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/ec2.html#waiters
waiter = client.get_waiter('instance_running')
waiter.wait(InstanceIds=[result['Instances'][0]['InstanceId']], WaiterConfig={
    'Delay': 3,
    'MaxAttempts': 5
})