#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/aws

terraform init
terraform plan
terraform apply -auto-approve

# Uncomment this to destroy the resources
# terraform destroy -auto-approve
# rm -rf .terraform
# rm -rf .terraform.lock.hcl
# rm -rf terraform.tfstate
# rm -rf terraform.tfstate.backup

cd ../../script/aws