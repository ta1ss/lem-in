#!/bin/sh
echo "Press any key to run the script..."
read -n 1 -s

echo "## RUNNING TEST WITH: example00 ##"
go run . example00.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example01 ##"
go run . example01.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example02 ##"
go run . example02.txt
echo -e "---------------------\n"

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example03 ##"
go run . example03.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example04 ##"
go run . example04.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example05 ##"
go run . example05.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: badexample00 ##"
go run . badexample00.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: badexample01 ##"
go run . badexample01.txt
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example06 ##"
start_time=$(date +%s)
go run . example06.txt
end_time=$(date +%s)
elapsed_time=$((end_time-start_time))
echo "Elapsed time: $elapsed_time seconds"
echo -e "---------------------\n" 

echo "Press any key to run next test"
read -n 1 -s

echo "## RUNNING TEST WITH: example07 ##"
start_time=$(date +%s)
go run . example07.txt
end_time=$(date +%s)
elapsed_time=$((end_time-start_time))
echo "Elapsed time: $elapsed_time seconds"
echo -e "---------------------\n" 
