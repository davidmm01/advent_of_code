#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>


int main(){

    std::ifstream input_file;  
    input_file.open("input.txt");

    std::string line_read;

    getline(input_file, line_read);

    for (int noun=0; noun <100; noun++){
        for (int verb=0; verb <100; verb++){

            std::stringstream ss(line_read);

            std::vector<int> input;

            while(ss.good()){
                std::string substring;
                getline(ss, substring, ',');
                input.push_back(stoi(substring));
            }

            int val;
            int cursor = 0;

            // before running the proigram replace the values at pos 1 with noun and 2 with verb
            input[1] = noun;
            input[2] = verb;

            while(true){
                if(input[cursor] == 99){
                    break;
                }
                else if(input[cursor] == 1){
                    val = input[input[cursor+1]] + input[input[cursor+2]];
                    input[input[cursor+3]] = val;
                }
                else if(input[cursor] == 2) {
                    val = input[input[cursor+1]] * input[input[cursor+2]];
                    input[input[cursor+3]] = val;
                }
                else {
                    std::cout << "unexpected opcode, exiting...";
                    return 0;
                }
                cursor+=4;
            }

            if (input[0] == 19690720){
                std::cout << "Solution: ";
                std::cout << 100 * noun + verb;
                return 0;
            }

        }
    }

    std::cout << "Error.  no solution found.";

}
