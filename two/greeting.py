from enum import Enum

class Greeting(Enum):
    WORKFLOW = "SayHelloWorkflow"
    MESSAGE = "from Two"
    QUEUE = "temporalab/one:worker"
