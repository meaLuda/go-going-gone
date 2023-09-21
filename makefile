proj1_test_build_image:
	docker build ./go-sqlite3/. 

proj2_test_build_image:
	docker build ./go-postgresql/dairy-api/. 

compose_img_db:
	docker-compose up -d --build