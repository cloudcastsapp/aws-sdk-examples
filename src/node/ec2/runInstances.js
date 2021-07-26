import { EC2Client, RunInstancesCommand, waitUntilInstanceRunning } from "@aws-sdk/client-ec2";

const ec2Client = new EC2Client({
    region: 'us-east-2',
    version: 'latest',
    // profile: 'your-profile',
})

const result = await ec2Client.send(new RunInstancesCommand({
    // ImageId: https://cloud-images.ubuntu.com/locator/ec2/
    ImageId: "ami-0b29b6e62f2343b46",
    InstanceType: "t3.small",

    // Optional parameters you likely want to define
    // KeyName: "some-key",
    // SecurityGroupIds: ["sg-foobar"],
    // SubnetId: "subnet-foobar",

    MinCount: 1,
    MaxCount: 1,
    BlockDeviceMappings: [{
        DeviceName: "/dev/sda1",
        Ebs: {
            DeleteOnTermination: true,
            VolumeType: "gp3",
            VolumeSize: 8
        }
    }],
}))

await waitUntilInstanceRunning({ ec2Client, maxWaitTime: 30, minDelay: 3 }, { InstanceIds: [result.Instances[0].InstanceId] })