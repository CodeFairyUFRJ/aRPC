namespace java calculator
namespace py calculator

enum Operation {
  ADD = 1,
  SUBTRACT = 2,
  MULTIPLY = 3,
  DIVIDE = 4
}

struct Work {
  1: i32 num1 = 0,
  2: i32 num2,
  3: Operation op,
  4: optional string comment,
}

exception InvalidOperation {
  1: i32 what,
  2: string why
}

service Calculator {

  void ping(),

  i32 add(1:i32 num1, 2:i32 num2),

  i32 calculate(1:i32 logid, 2:Work w) throws (1:InvalidOperation ouch),

  /**
   * Este metodo tem o modificador oneway, 
   * que implica que o cliente apenas faz a requisicao mas
   * nao espera a resposta.
   * Metodos com oneway devem retornar void.
   */
  oneway void zip()
}