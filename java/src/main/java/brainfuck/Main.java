package brainfuck;

import java.io.FileReader;
import java.io.Reader;
import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

public class Main {
    public static Token[] compile(Reader reader) throws Exception {
        List<Token> tokens = new ArrayList<>();
        Stack<Integer> stack = new Stack<>();

        int pos = -1;
        while (true) {
           int b = reader.read();
           if (b == -1) {
               if (!stack.empty()) throw new Exception(String.format("no corresponding ] for [ at pos %d", stack.pop()));
               break;
           }
           pos++;

           switch (b) {
               case '>':
                   tokens.add(new Token('>', -1));
                   break;
               case '<':
                   tokens.add(new Token('<', -1));
                   break;
               case '+':
                   tokens.add(new Token('+', -1));
                   break;
               case '-':
                   tokens.add(new Token('-', -1));
                   break;
               case '.':
                   tokens.add(new Token('.', -1));
                   break;
               case ',':
                   tokens.add(new Token(',', -1));
                   break;
               case '[':
                   tokens.add(new Token('[', -1));
                   stack.push(pos);
                   break;
               case ']':
                   if (stack.empty()) throw new Exception(String.format("no corresponding [ for ] at pos %d", pos));
                   int pos1 = stack.pop();
                   Token token1 = new Token('[', pos);
                   tokens.set(pos1, token1);
                   tokens.add(new Token(']', pos1));
                   break;
               default:
                   throw new Exception(String.format("invalid operator %c", b));
           }
        }

        return tokens.toArray(new Token[tokens.size()]);
    }

    public static void vm(Token[] tokens) {
        int[] memory = new int[3000];
        int index = 0;

        int pos = 0;
        while (true) {
            if (pos == tokens.length) break;
            Token token = tokens[pos];

            switch (token.getOp()) {
                case '>':
                    index++;
                    pos++;
                    break;
                case '<':
                    index--;
                    pos++;
                    break;
                case '+':
                    memory[index]++;
                    pos++;
                    break;
                case '-':
                    memory[index]--;
                    pos++;
                    break;
                case '.':
                    System.out.printf("%c", memory[index]);
                    pos++;
                    break;
                case ',':
                    // not implement
                    break;
                case '[':
                    if (memory[index] == 0) pos = token.getPos();
                    pos++;
                    break;
                case ']':
                    pos = token.getPos();
                    break;
                default:
                    // do nothing
            }
        }
    }

    public static void main(String[] args) throws Exception {
        Reader reader;
        Token[] tokens;

        reader = new FileReader("../test/integer.bf");
        tokens = compile(reader);
        vm(tokens);
        reader.close();
        System.out.println();

        reader = new FileReader("../test/cycle.bf");
        tokens = compile(reader);
        vm(tokens);
        reader.close();
        System.out.println();

        reader = new FileReader("../test/helloworld.bf");
        tokens = compile(reader);
        vm(tokens);
        reader.close();
    }
}
