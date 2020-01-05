#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <iomanip>


std::vector<int> get_input_vector(std::string filename){

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
    return input;
}


std::string pad_instruction(int raw){
    // takes an integer and pads out missing leading digits with zeroes,
    // creating a string that always has 5 places
    std::stringstream ss;
    ss << std::setw(5) << std::setfill('0') << raw;
    return ss.str();
}


int determine_cursor_delta(std::string instruction){
    if (instruction[5] == '1' or instruction[5] == '2'){
        return 4;
    }
    else if (instruction[5] == '3' or instruction[5] == '4') {
        return 1;
    }
    else {
        std::cout << "Error determining cursor delta\n";
        std::cout << "Unexpected value: ";
        std::cout << instruction[5];
    }
}


int main(){

    std::vector<int> input = get_input_vector("input.txt");

    int val;
    int cursor = 0;

    // for(std::size_t i=0; i<input.size(); i++){
    //     std::cout << pad_instruction(input[0]) << std::endl;

    // }


    while(true){

        int raw_instruction = input[cursor];

        std::string padded_instruction = pad_instruction(raw_instruction);



        // // opcode 01: ADDITION PROTOCOL
        // if(input[cursor] == 1){
        //     val = input[input[cursor+1]] + input[input[cursor+2]];
        //     input[input[cursor+3]] = val;
        // }

        // // opcode 02: MULTIPLICATION PROTOCOL
        // else if(input[cursor] == 2) {
        //     val = input[input[cursor+1]] * input[input[cursor+2]];
        //     input[input[cursor+3]] = val;
        // }

        // // opcode 03: SAVE PROTOCOL


        // // opcode 04: OUTPUT PROTOCOL


        // // opcode 99: HALT PROGRAM
        // else if(input[cursor] == 99){
        //     break;
        // }


        // else {
        //     std::cout << "unexpected opcode, exiting...";
        //     return 0;
        // }
        
        cursor += determine_cursor_delta(padded_instruction);
    
    }





    return 0;

}
