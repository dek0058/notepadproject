#include <iostream>
#include <vector>
#include "quicksort.h"
#include "mergesort.h"

int main() {
    
    std::vector<int> arr = {10, 7, 8, 9, 1, 5, 12};
    std::cout << "Original array: ";
    for (const auto &num : arr) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    sort::MergeSort<int>(arr, std::greater());
    std::cout << "Sorted array: ";
    for (const auto &num : arr) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    return 0;
}
