import json
import boto3

def lambda_handler(event, context):
    s3 = boto3.resource('s3')
    code_pipeline = boto3.client('codepipeline')
    
    bucket = s3.Bucket('thehotbox.xyz')
    
    objs = bucket.objects.filter(Prefix="frontend/")
    
    for o in objs:
        s3.Object('thehotbox.xyz', remove_parent_directory(o.key)).copy_from(CopySource='thehotbox.xyz/'+o.key)
        s3.Object('thehotbox.xyz', o.key).delete()
        
    
    code_pipeline.put_job_success_result(jobId=event['CodePipeline.job']['id'])
    
    return {
        'statusCode': 200,
        'body': 'success'
    }

def remove_parent_directory(s):
    return s[s.index('frontend/') + 9:]
