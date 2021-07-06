import { S3Client, CreateBucketCommand, waitUntilBucketExists } from "@aws-sdk/client-s3";

const s3Client = new S3Client({
    region: 'us-east-2',
    version: 'latest',
    // profile: 'your-profile',
});

// Search for CreateBucketCommandInput to see input
// to pass to CreateBucketCommand
const Bucket = 'bucket-o-fun'

// Create the Amazon S3 bucket

// data: {Location: "us-east-2"}
// https://docs.aws.amazon.com/zh_cn/AWSJavaScriptSDK/v3/latest/clients/client-s3-control/classes/createbucketcommand.html
const data = await s3Client.send(new CreateBucketCommand({
    Bucket: Bucket
}));

// https://docs.aws.amazon.com/AWSJavaScriptSDK/v3/latest/clients/client-s3/globals.html#waituntilbucketexists
await waitUntilBucketExists({ s3Client, maxWaitTime: 30, minDelay: 3 }, { Bucket })

