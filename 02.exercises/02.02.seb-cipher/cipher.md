# Seb Cipher

The name is provisional, because I do not know if a similar algorithm already exists — most likely it does. It was inspired by the rail fence cipher.

1. The first letter of the message is added, followed by random characters. The number of random characters is given by the key.
2. The second letter of the message is added, followed by random characters. The number of random characters is one less than the number of random characters used previously.
3. This repeats until zero random characters have been added after a letter. From this point, an ascending number of random characters is added to following letters.
4. This repeats until the number of random characters added after a letter is equal to the key. At this point, the process repeats.
5. The process ends when the end of the message has been reached.

It should be noted that there are no safeguards by default to prevent unfortunate combinations of random letters from being generated.

## Example

- Message: CROCODILE
- Key: 3

"%" denotes a random character that will be generated during runtime.

- C % % %
- R % %
- O %
- C
- O %
- D % %
- I % % %
- L % %
- E %

- Result: C%%%R%%O%CO%D%%I%%%L%%E%
- Result: **C**UBB**R**SU**O**U**CO**L**D**TV**I**GIF**L**ZF**E**S