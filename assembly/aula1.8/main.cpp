#include <iostream>

extern "C" int GetValueASM(int x);

int main() {
  std::cout << "From Assembly code: " << GetValueASM(32) << "\n";

  return 0;
}
