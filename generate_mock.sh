echo "==generating mockfile for repository=="
mockgen -source=./internal/repository/cake.go -destination=./mock/repository/cake.go
echo "==mockfile for repository generated=="
echo "==generating mockfile for api handler=="
mockgen -source=./internal/api/cake.go -destination=./mock/api/cake.go
echo "==mockfile for api handler generated==" 
