#include <iostream>
#include <fstream>

int main(){

    std::ifstream input_file;  
    input_file.open("input.txt");

    std::string line_read;

    int fuel_required = 0;

    while(getline(input_file, line_read)){

            if (line_read == "\n"){
                continue;
            }

            int mass = stoi(line_read);
            int fuel = (mass / 3) - 2;
            fuel_required += fuel;

            // part 2
            while (fuel > 0) {
                
                /*
                    interesting GOTCHA!
                    before this solution i had 
                        int fuel = (fuel / 3) - 2;
                    So both using and reassigning over the fuel variable (fuels_fuel did not exist, also fuels_fuel is a TERRIBLE variable name)
                    This didn't work.  The value of fuel would never change and i got stuck in an infinite loop.
                    TODO: find out why!
                */

                int fuels_fuel = (fuel / 3) - 2;

                if (fuels_fuel > 0){
                    fuel_required += fuels_fuel;
                }
                
                fuel = fuels_fuel;
            }

    }

    std::cout << "total fuel required: ";
    std::cout << fuel_required << std::endl;
    return 0;

}
