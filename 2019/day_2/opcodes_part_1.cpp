#include <iostream>
#include <fstream>
#include <sstream>  // gives us the stringstream class
#include <vector>  // gives us the vector type


int main(){

    std::ifstream input_file;  
    input_file.open("input.txt");

    std::string line_read;

    // input only has 1 line
    getline(input_file, line_read);

    std::stringstream ss(line_read);

    // use vector instead of an array to have variable size support
    std::vector<int> input;

    while(ss.good()){
        std::string substring;
        getline(ss, substring, ',');
        input.push_back(stoi(substring));
    }

    int val;
    int cursor = 0;

    // before running the proigram replace the values at pos 1 with 12 and 2 with 2
    input[1] = 12;
    input[2] = 2;

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

    std::cout << "Solution: ";
    std::cout << input[0] << std::endl;
    
    return 0;

}
