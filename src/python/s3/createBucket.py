import boto3

bucket_name = 'bucket-o-fun'

# Create S3 client using default configuration
# s3 = boto3.resource('s3')
# Create S3 client defining some options
# See https://boto3.amazonaws.com/v1/documentation/api/latest/reference/core/session.html
# See https://stackoverflow.com/questions/33378422/how-to-choose-an-aws-profile-when-using-boto3-to-connect-to-cloudfront
ses = boto3.session.Session(region_name='us-east-2') # profile_name='your-profile'
client = ses.resource('s3')

# See https://boto3.amazonaws.com/v1/documentation/api/latest/guide/s3-example-creating-buckets.html
location = {'LocationConstraint': 'us-east-2'}
client.create_bucket(Bucket=bucket_name, CreateBucketConfiguration=location)

# See https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/s3.html#S3.Waiter.BucketExists
waiter = client.get_waiter('bucket_exists')
waiter.wait(Bucket=bucket_name, WaiterConfig={
    'Delay': 3,
    'MaxAttempts': 5
})