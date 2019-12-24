//
// Created by czx on 2019/12/23.
//

#ifndef DATASTRUCT_ALGORITEM_CPLUSPLUS_VECTOR_HPP
#define DATASTRUCT_ALGORITEM_CPLUSPLUS_VECTOR_HPP

template<typename E>
class Vector {
public:
    typedef E *iterator;
    typedef const E *const_iterator;
    enum {
        SPACE_CAPACITY = 16
    };
public:
    explicit Vector(int initSize = 0) : length(initSize), arrayCapacity(initSize + SPACE_CAPACITY) {
        element = new int[arrayCapacity];
    }

    Vector(const Vector &rhs) : element(nullptr) {
        operator=(rhs);
    }

    ~Vector() { delete element; }

    const Vector &operator=(const Vector &rhs) {
        if (this != rhs) {
            delete[] element;
            length = rhs.length;
            arrayCapacity = rhs.arrayCapacity;
            element = new E[arrayCapacity];
            std::copy(rhs, rhs + length, element);
        }
        return *this;
    }

    void resize(int newSize) {
        if (newSize > arrayCapacity) {
            reserve(newSize * 2);
        }
        length = newSize;
    }

    void reserve(int newCapacity) {
        if (newCapacity < length)
            return;

        E *old = element;
        element = new E[newCapacity];
        std::copy(old, old + length, element);
        arrayCapacity = newCapacity;
        delete[] old;
    }

    E &operator[](int index) const {
        if (index > length) {
            throw std::runtime_error("index out bound");
        }
        return element[index];
    }

    bool empty() const {
        return length == 0;
    }

    int size() const { return length; }

    int capacity() const { return arrayCapacity; }

    void push_back(const E &e) {
        if (length == arrayCapacity) {
            reserve(2 * arrayCapacity);
        }
        element[length++] = e;
    }

    void pop_back() { --length; }

    E &back() {
        return element[--length];
    }

    void insert(int index, const E &e) {
        if (index < 0 || index > length)
            return;
        if (length == arrayCapacity) {
            reserve(2 * arrayCapacity);
        }
        // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
        std::copy_backward(element + index, element + length, element + length+1);
        element[index] = e;
        ++length;
    }

    iterator begin() {
        return &element[0];
    }

    const_iterator begin() const { return &element[0]; }

    iterator end() {
        return &element[length];
    }

    const_iterator end(const) {
        return &element[length];
    }

private:
    E *element;
    int length;
    int arrayCapacity;
};

#endif //DATASTRUCT_ALGORITEM_CPLUSPLUS_VECTOR_HPP
