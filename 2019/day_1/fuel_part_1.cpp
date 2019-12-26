#include <iostream>  // for stream I/O, e.g. printing our solution to the terminal
#include <fstream>  // stream class to both read and write from/to files

// main function is the entry point of cpp programs
int main(){

    std::ifstream input_file;  
    // ifstream class is to read (only) from files, rather than fstream or ofstream
    // it is just setting the mode parameter to ios::in rather than ios:out (writing) or a combination of the two
    input_file.open("input.txt");

    std::string line_read;

    int fuel_required = 0;

    while(getline(input_file, line_read)){

            if (line_read == "\n"){
                // deal with the last line of the input file
                continue;
            }

            int mass = stoi(line_read);  // convert string to an int for arithmetic

            // integer division is always rounded down
            int fuel = (mass / 3) - 2;
            fuel_required += fuel;
    }

    std::cout << "total fuel required: ";
    std::cout << fuel_required << std::endl;
    return 0;

}
