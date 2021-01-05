from locust import HttpUser, between, task


class WebsiteUser(HttpUser):
    wait_time = between(5, 15)

    @task
    def data(self):
        for i in range(100):
            self.client.get("/?id="+str(i))
