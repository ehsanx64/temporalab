import asyncio
import uuid
from temporalio.client import Client 
from greeting import Greeting

async def main():
    client = await Client.connect("localhost:7233")
    result = await client.execute_workflow(
            Greeting.WORKFLOW.value, 
            Greeting.MESSAGE.value,
            id=f"two::start:{uuid.uuid4()}",
            task_queue=Greeting.QUEUE.value,
    )

    print("Workflow result: ", result)

if __name__ == "__main__":
    asyncio.run(main())
