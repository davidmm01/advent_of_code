#include <iostream>
#include <vector>
#include <sstream>


// our puzzle inputs
int LOWER_BOUND = 234208;
int UPPER_BOUND = 765869;


bool is_only_increasing(int password){

    // convert integer to a string so we can iterate over each individual element
    std::string digits = std::to_string(password);
    
    int previous = 0;  // this is safe since the first digit will never be a 0
    for(char& digit : digits) {
        int value = digit - '0';
        if (value < previous){
            return false;
        }
        else {
            previous = value;
        }
    }
    return true;
}


bool has_strictly_double_digit(int password){

    std::vector<int> digit_count;    


    std::string digits = std::to_string(password);

    int previous = 0;  // this is safe since the first digit will never be a 0
    int count = 0;

    for(char& digit : digits) {

        int value = digit - '0';

        if (value == previous){
            count ++;
        }
        else {
            digit_count.push_back(count);
            previous = value;
            count = 1;
        }

    }

    // if the final number was part of a group, it has not been pushed yet
    if(count > 1){
        digit_count.push_back(count);
    }

    for(std::size_t i=0; i<digit_count.size(); i++){
        if (digit_count[i] == 2){
            return true;
        }
    }
    return false;
}


bool do_tests_pass(){
    if (is_only_increasing(223450)){
        std::cout << "LOGIC FAIL: is_only_increasing should have returned FALSE for '223450'\n";
        return false;
    }

    if (!is_only_increasing(223456)){
        std::cout << "LOGIC FAIL: is_only_increasing should have returned TRUE for '223456'\n";
        return false;
    }

    if (has_strictly_double_digit(345678)){
        std::cout << "LOGIC FAIL: has_strictly_double_digit should have returned FALSE for '345678'\n";
        return false;
    }

    if (has_strictly_double_digit(123444)){
        std::cout << "LOGIC FAIL: has_strictly_double_digit should have returned FALSE for '123444'\n";
        return false;
    }

    if (!has_strictly_double_digit(111122)){
        std::cout << "LOGIC FAIL: has_strictly_double_digit should have returned TRUE for '111122'\n";
        return false;
    }

    return true;

}


int main(){

    if(!do_tests_pass()){
        std::cout << "Failed tests... exiting\n";
        return -1;
    }

    std::vector<int> password_candidates;    

    for(int i = LOWER_BOUND; i<=UPPER_BOUND; i++){
        if(is_only_increasing(i) and has_strictly_double_digit(i)){
            password_candidates.push_back(i);
        }
    }

    int num_candidates = password_candidates.size(); 

    std::cout << "Solution: ";
    std::cout << num_candidates;
    return 0;

}
