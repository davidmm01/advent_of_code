#include <iostream>
#include <fstream>
#include <sstream>  // gives us the stringstream class
#include <vector>  // gives us the vector type


int main(){

    std::ifstream input_file;  
    input_file.open("input.txt");

    std::string line_read;

    // initialise a 20000*20000 matrix of all zeroes, and some cursors for navigating it
    // to work for different inputs, it might be necessary to make the wire_grid bigger, which is a flaw of the design
    std::vector<std::vector<int>> wire_grid(20000, std::vector<int>(20000));
    int x_cursor = 10000;
    int y_cursor = 10000;
    // target variables will track where the cursors should be moving towards
    int x_target = 10000;
    int y_target = 10000;

    // in hindsight, this approach is really wasteful, there is a good chance that most of the matrix is just full  of zeroes.
    // would be better to use a key&value lookup to track where has been navigated to so far, but with less zeroes everywhere.
    // in python id use a dictionary of dictionaries.

    // also in hindsight, hardcoding all these dimensions is so bad.  better to iterate through input and determine an appropriately sized
    // matrix. then calculate midpoints etc from that.  but we are aiming for progress not perfection right now :P

    // grab the first line of the input file and break it into an iterable vector
    getline(input_file, line_read);
    std::stringstream ss(line_read);
    std::vector<std::string> wire_1_path;
    while(ss.good()){
        std::string substring;
        getline(ss, substring, ',');
        wire_1_path.push_back(substring);
    }

    // iterate over the wire path, filling in the wire_grid with 1's where it goes
    for(std::size_t i=0; i<wire_1_path.size(); i++){

        char direction = wire_1_path[i].at(0);


        int magnitude = stoi(wire_1_path[i].substr(1));  // leaving off 2 arg (length) returns whole remaining string

        if (direction == 'L') {
            x_target = x_cursor - magnitude;
            for (x_cursor; x_cursor > x_target; x_cursor--){
                wire_grid[x_cursor][y_cursor] = 1;
            }
        }
        else if (direction == 'R') {
            x_target = x_cursor + magnitude;
            for (x_cursor; x_cursor < x_target; x_cursor++){
                wire_grid[x_cursor][y_cursor] = 1;
            }
        }
        else if (direction == 'U') {
            y_target = y_cursor + magnitude;
            for (y_cursor; y_cursor < y_target; y_cursor++){
                wire_grid[x_cursor][y_cursor] = 1;
            }
        }
        else if (direction == 'D') {
            y_target = y_cursor - magnitude;
            for (y_cursor; y_cursor > y_target; y_cursor--){
                wire_grid[x_cursor][y_cursor] = 1;
            }
        }

    }

    // grab the second line of input and build its path
    getline(input_file, line_read);
    std::stringstream ss2(line_read);
    std::vector<std::string> wire_2_path;
    while(ss2.good()){
        std::string substring;
        getline(ss2, substring, ',');
        wire_2_path.push_back(substring);
    }

    // reset cursors
    x_cursor = 10000;
    y_cursor = 10000;
    x_target = 10000;
    y_target = 10000;

    // create vector for storing intersections
    std::vector<std::vector<int>> intersections;

    // lots of duplication :(

    for(std::size_t i=0; i<wire_2_path.size(); i++){
        char direction = wire_2_path[i].at(0);
        int magnitude = stoi(wire_2_path[i].substr(1));

        if (direction == 'L') {
            x_target = x_cursor - magnitude;
            for (x_cursor; x_cursor > x_target; x_cursor--){
                if (wire_grid[x_cursor][y_cursor] == 1){
                    intersections.push_back(std::vector<int>{x_cursor, y_cursor});
                }
            }
        }
        else if (direction == 'R') {
            x_target = x_cursor + magnitude;
            for (x_cursor; x_cursor < x_target; x_cursor++){
                if (wire_grid[x_cursor][y_cursor] == 1){
                    intersections.push_back(std::vector<int>{x_cursor, y_cursor});

                }            
            }
        }
        else if (direction == 'U') {
            y_target = y_cursor + magnitude;
            for (y_cursor; y_cursor < y_target; y_cursor++){
                if (wire_grid[x_cursor][y_cursor] == 1){
                    intersections.push_back(std::vector<int>{x_cursor, y_cursor});

                }            
            }
        }
        else if (direction == 'D') {
            y_target = y_cursor - magnitude;
            for (y_cursor; y_cursor > y_target; y_cursor--){
                if (wire_grid[x_cursor][y_cursor] == 1){
                    intersections.push_back(std::vector<int>{x_cursor, y_cursor});
                }            
            }
        }

    }

    int smallest_distance = 1410065408;
    int distance;

    for(std::size_t i=0; i<intersections.size(); i++){
        distance = abs(10000-intersections[i][0]) + abs(10000-intersections[i][1]);
        if (distance < smallest_distance and distance != 0){
            smallest_distance = distance;
        }
    }

    std::cout << "Solution: ";
    std::cout << smallest_distance;

    return 0;

}
