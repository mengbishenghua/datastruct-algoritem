//
// Created by czx on 2019/12/23.
//
#include <iostream>
#include "Vector.hpp"
#include "List.hpp"

void vectorTest();

using namespace std;

void vectorTest() {
    Vector<int> v;
    for (int i = 0; i < 10; ++i) {
        v.push_back(i);
    }

    for (int &i : v) {
        cout << i << " ";
    }

    v.insert(5, 100);
    cout << endl;
    for (auto it = v.begin(); it != v.end(); ++it) {
        cout << *it << " ";
    }
}

void listTest() {
    List<int> list;
    for (int i = 0; i < 10; ++i) {
        list.push_back(i);
    }

    for (int &i : list) {
        cout << i << " ";
    }
}

int main(int argv, char *args[]) {
    vectorTest();
    std::cout << "\n---------------------------" << std::endl;
    listTest();
    return 0;
}
