//
// Created by czx on 2019/12/23.
//

using namespace std;

int main(int argv, char *args[]) {
    auto *x = new ArrayList<int>(100);
    ArrayList<double> y(100);

    ArrayList<double> w(y);

    for (int i = 0; i < 10; ++i) {
        x->insert(x->size(), i);
    }

    std::cout << x << std::endl;

    return 0;
}