from locust import task, HttpUser

class HelloWorldUser(HttpUser):
	@task
	def go1(self):
 		self.client.get("/go/sha256")
	@task
	def node1(self):
		self.client.get("/node/sha256")
