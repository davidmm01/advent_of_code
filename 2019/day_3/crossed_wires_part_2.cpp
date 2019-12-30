#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>


int main(){

    std::ifstream input_file;  
    input_file.open("input.txt");

    std::string line_read;

    std::vector<std::vector<int>> wire_grid(20000, std::vector<int>(20000));
    int x_cursor = 10000;
    int y_cursor = 10000;
    int x_target = 10000;
    int y_target = 10000;

    getline(input_file, line_read);
    std::stringstream ss(line_read);
    std::vector<std::string> wire_1_path;
    while(ss.good()){
        std::string substring;
        getline(ss, substring, ',');
        wire_1_path.push_back(substring);
    }

    int wire_1_steps = -1;  // start at -1 since when we reach an intersection, we wont all the steps leading up to that intersection

    // iterate over the wire path, filling in the wire_grid with 1's where it goes
    for(std::size_t i=0; i<wire_1_path.size(); i++){

        char direction = wire_1_path[i].at(0);


        int magnitude = stoi(wire_1_path[i].substr(1));

        if (direction == 'L') {
            x_target = x_cursor - magnitude;
            for (x_cursor; x_cursor > x_target; x_cursor--){
                wire_1_steps++;
                if (wire_grid[x_cursor][y_cursor] == 0 and !(x_cursor == 10000 and y_cursor == 10000)){
                    // only record the step count if its the first time the point has been visited
                    wire_grid[x_cursor][y_cursor] = wire_1_steps;
                }
            }
        }
        else if (direction == 'R') {
            x_target = x_cursor + magnitude;
            for (x_cursor; x_cursor < x_target; x_cursor++){
                wire_1_steps++;
                if (wire_grid[x_cursor][y_cursor] == 0 and !(x_cursor == 10000 and y_cursor == 10000)){
                    // only record the step count if its the first time the point has been visited
                    wire_grid[x_cursor][y_cursor] = wire_1_steps;
                }
            }
        }
        else if (direction == 'U') {
            y_target = y_cursor + magnitude;
            for (y_cursor; y_cursor < y_target; y_cursor++){
                wire_1_steps++;
                if (wire_grid[x_cursor][y_cursor] == 0 and !(x_cursor == 10000 and y_cursor == 10000)){
                    // only record the step count if its the first time the point has been visited
                    wire_grid[x_cursor][y_cursor] = wire_1_steps;
                }
            }
        }
        else if (direction == 'D') {
            y_target = y_cursor - magnitude;
            for (y_cursor; y_cursor > y_target; y_cursor--){
                wire_1_steps++;
                if (wire_grid[x_cursor][y_cursor] == 0 and !(x_cursor == 10000 and y_cursor == 10000)){
                    // only record the step count if its the first time the point has been visited
                    wire_grid[x_cursor][y_cursor] = wire_1_steps;
                }
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


    int wire_2_steps = -1;  // start at -1 since when we reach an intersection, we wont all the steps leading up to that intersection

    // part 2 extension: here we will not only store the coords of the intersection, but also the steps up to that point for wire 2 
    // (only the lowest though, so we need to go through our list of intersections and each time...)
    for(std::size_t i=0; i<wire_2_path.size(); i++){

        char direction = wire_2_path[i].at(0);
        int magnitude = stoi(wire_2_path[i].substr(1));

        if (direction == 'L') {
            x_target = x_cursor - magnitude;
            for (x_cursor; x_cursor > x_target; x_cursor--){
                wire_2_steps++;
                if (wire_grid[x_cursor][y_cursor] != 0){
                    int found = 0;
                    for(std::size_t i=0; i<intersections.size(); i++){
                        if (intersections[i][0]==x_cursor and intersections[i][1]==y_cursor){
                            found = 1;
                            break;
                        }
                    }
                    // dont store intersection if one already exists, since any already existing intersection will have lower steps
                    if (found == 0){
                        intersections.push_back(std::vector<int>{x_cursor, y_cursor, wire_2_steps});
                    }
                }
            }
        }
        else if (direction == 'R') {
            x_target = x_cursor + magnitude;
            for (x_cursor; x_cursor < x_target; x_cursor++){
                wire_2_steps++;
                if (wire_grid[x_cursor][y_cursor] != 0){
                    int found = 0;
                    for(std::size_t i=0; i<intersections.size(); i++){
                        if (intersections[i][0]==x_cursor and intersections[i][1]==y_cursor){
                            found = 1;
                            break;
                        }
                    }
                    // dont store intersection if one already exists, since any already existing intersection will have lower steps
                    if (found == 0){
                        intersections.push_back(std::vector<int>{x_cursor, y_cursor, wire_2_steps});
                    }
                }            
            }
        }
        else if (direction == 'U') {
            y_target = y_cursor + magnitude;
            for (y_cursor; y_cursor < y_target; y_cursor++){
                wire_2_steps++;
                if (wire_grid[x_cursor][y_cursor] != 0){
                    int found = 0;
                    for(std::size_t i=0; i<intersections.size(); i++){
                        if (intersections[i][0]==x_cursor and intersections[i][1]==y_cursor){
                            found = 1;
                            break;
                        }
                    }
                    // dont store intersection if one already exists, since any already existing intersection will have lower steps
                    if (found == 0){
                        intersections.push_back(std::vector<int>{x_cursor, y_cursor, wire_2_steps});
                    }
                }            
            }
        }
        else if (direction == 'D') {
            y_target = y_cursor - magnitude;
            for (y_cursor; y_cursor > y_target; y_cursor--){
                wire_2_steps++;
                if (wire_grid[x_cursor][y_cursor] != 0){
                    int found = 0;
                    for(std::size_t i=0; i<intersections.size(); i++){
                        if (intersections[i][0]==x_cursor and intersections[i][1]==y_cursor){
                            found = 1;
                            break;
                        }
                    }
                    // dont store intersection if one already exists, since any already existing intersection will have lower steps
                    if (found == 0){
                        intersections.push_back(std::vector<int>{x_cursor, y_cursor, wire_2_steps});
                    }
                }            
            }
        }

    }

    int smallest_distance = 1410065408;
    int distance;

    for(std::size_t i=0; i<intersections.size(); i++){
        // wire 2 steps is 3rd ele of intersections array
        // wire 1 steps is the number in the big matrix.  add these
        distance = intersections[i][2] + wire_grid[intersections[i][0]][intersections[i][1]];
        if (distance < smallest_distance and distance != 0){
            smallest_distance = distance;
        }
    }

    std::cout << "Solution: ";
    std::cout << smallest_distance;

    return 0;

}
