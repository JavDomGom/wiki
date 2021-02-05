import aioboto3

async def main():
    async with aioboto3.resource("s3") as s3:
        bucket = await s3.Bucket('mybucket')  # <----------------
        async for s3_object in bucket.objects.all():
            print(s3_object)