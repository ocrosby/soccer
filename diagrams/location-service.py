from diagrams import Diagram, Cluster
from diagrams.aws.database import RDS
from diagrams.aws.database import ElastiCache
from diagrams.aws.compute import EKS
from diagrams.aws.compute import ECS  # Assuming location-service runs on ECS for illustration

with Diagram("Location Service Architecture", show=False, direction='TB', filename="../cmd/location-service/assets/location_service_architecture"):
    with Cluster("EKS Cluster"):
        redis = ElastiCache("Redis Cache")
        postgres = RDS("PostgreSQL DB")
        location_service = ECS("Location Service")

        location_service >> redis >> postgres
