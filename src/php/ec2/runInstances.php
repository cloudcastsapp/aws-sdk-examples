<?php

require_once(__DIR__.'/../vendor/autoload.php');

use Aws\Ec2\Ec2Client;

$ec2 = new Ec2Client([
    'region' => 'us-east-2',
    'version' => 'latest',
    // 'profile' => 'your-profile',
]);

// $result: ['Location' => 'us-east-2']
$result = $ec2->runInstances([
    // ImageId: https://cloud-images.ubuntu.com/locator/ec2/
    'ImageId' => 'ami-0b29b6e62f2343b46',
    'InstanceType' => 't3.small',

    // Optional parameters you likely want to define
    // 'KeyName' => "some-key",
    // 'SecurityGroupIds' => ['sg-foobar'],
    // 'SubnetId' => 'subnet-foobar',

    'MaxCount' => 1,
    'MinCount' => 1,
    'BlockDeviceMappings' => [
        'DeviceName' => '/dev/sda1',
        'Ebs' => [
            'DeleteOnTermination' => true,
            'VolumeSize' => 8,
            'VolumeType' => 'gp3',
        ],
    ],
]);

$ec2->waitUntil('InstanceRunning', [
    'InstanceIds'  => [$result['Instances'][0]['InstanceId']],
    '@waiter' => [
        'delay'       => 3, // Wait 3 seconds between polling
        'maxAttempts' => 10, // Max attempts before failing, total of 30 seconds waited here
    ]
]);