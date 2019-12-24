//
// Created by czx on 2019/12/24.
//

#ifndef DATASTRUCT_ALGORITEM_CPLUSPLUS_LIST_HPP
#define DATASTRUCT_ALGORITEM_CPLUSPLUS_LIST_HPP

#include <exception>
#include <stdexcept>

template<typename E>
class List {
private:
    struct Node {
        Node *prev;
        Node *next;
        E data;

        explicit Node(Node *prev = nullptr, Node *next = nullptr, const E &data = E())
                : prev(prev), next(next), data(data) {}
    };

public:
    class const_iterator {
    public:
        const_iterator() : current(nullptr) {}

        const E &operator*() const {
            return retrieve();
        }

        const_iterator &operator++() {
            current = current->next;
            return *this;
        }

        const_iterator operator++(int) {
            const_iterator old = *this;
            ++(*this);
            return old;
        }

        bool operator==(const const_iterator &rhs) const {
            return current == rhs.current;
        }

        bool operator!=(const const_iterator &rhs) const {
            return !(*this == rhs);
        }

    public:
        const List<E> *thenList;
        Node *current;

        E &retrieve() const {
            return current->data;
        }

        explicit const_iterator(const List<E> &thenList, Node *p) : thenList(&thenList), current(p) {}

        void assertIsValid() const {
            if (thenList == nullptr || current == nullptr || current == thenList->head) {
                throw std::runtime_error("iterator out of bound exception");
            }
        }

        friend class List<E>;
    };

    class iterator : public const_iterator {
    public:
        iterator() = default;

        E &operator*() {
            return List::const_iterator::retrieve();
        }

        const E &operator*() const {
            return const_iterator::operator*();
        }

        iterator &operator++() {
            List::const_iterator::current = List::const_iterator::current->next;
            return *this;
        }

        iterator operator++(int) {
            iterator old = *this;
            ++(*this);
            return old;
        }

    protected:
        explicit iterator(const List<E> &thenList, Node *p) : const_iterator(thenList, p) {}

        friend class List<E>;

    };

public:
    List() { init(); }

    ~List() {
        clear();
        delete head;
        delete tail;
    }

    List(const List &rhs) {
        init();
        *this = rhs;
    }

    const List &operator=(const List &rhs) {
        if (*this == rhs)
            return *this;
        clear();
        for (const_iterator it = rhs.begin(); it != rhs.end(); ++it) {
            push_back(*it);
        }
        return *this;
    }

    iterator begin() {
        iterator it(*this, head);
        return ++it;
    }

    iterator end() {
        iterator it(*this, tail);
        return it;
    }

    const_iterator begin() const {
        const_iterator it(*this, head);
        return ++it;
    }

    const_iterator end() const {
        const_iterator it(*this, tail);
        return it;
    }

    int size() { return thenSize; }

    bool empty() { return thenSize == 0; }

    void clear() {
        while (!empty()) {
            pop_front();
        }
    }

    E &front() {
        return *begin();
    }

    const E &front() const {
        return *begin();
    }

    // 先--后*
    E &back() { return *--end(); }

    const E &back() const { return *--end(); }

    void push_front(const E &e) {
        insert(begin(), e);
    }

    void push_back(const E &e) {
        insert(end(), e);
    }

    void pop_front() {
        erase(begin());
    }

    void pop_back() {
        erase(--end());
    }

    iterator insert(iterator it, const E &e) {
        it.assertIsValid();
        if (it.thenList != this) {
            throw std::runtime_error("iterator Mismatch exception");
        }
        Node *p = it.current;
        // 拆为3步
        // 1. newNode = new Node(p->prev, p, e)
        // 2.p->prev->next = newNode
        // 3.p->prev = newNode;
        p->prev = p->prev->next = new Node(p->prev, p, e);
        ++thenSize;
        // 返回插入的节点位置
        return iterator(*this, p->prev);
    }

    iterator erase(iterator it) {
        Node *p = it.current;
        iterator retVal(*this, p->next);
        p->prev->next = p->next;
        p->next->prev = p->prev;
        delete p;
        --thenSize;
        // 返回删除元素的后一个元素
        return retVal;
    }

    iterator erase(iterator start, iterator end) {
        for (iterator it = start; it != end;) {
            // 删除元素后，it失效了，要重新赋值为删除元素的后一个元素
            it = erase(it);
        }
    }

private:
    int thenSize;
    Node *head;
    Node *tail;

    void init() {
        thenSize = 0;
        head = new Node;
        tail = new Node;
        head->next = tail;
        tail->prev = head;
    }
};

#endif //DATASTRUCT_ALGORITEM_CPLUSPLUS_LIST_HPP
