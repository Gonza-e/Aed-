Program Ejercicio1;

{1) Escribir un programa en pascal que dado dos números enteros realice la suma, resta,
multiplicación, división entera y resto de los mismos. Los numero son: X = 55 e Y = 30.}

var 
	suma, mul, division, resto, resta: integer;
const 
	x = 55; 
	y = 30;
begin

	suma:= x + y;
	division:= x div y; 
	mul:= x * y;
	resto:= x mod y; 
	resta:= x - y;

	writeln('Suma: ',suma);
	writeln('Division: ',division);
	writeln('Multiplicacion: ',mul);
	writeln('Resto: ',resto);
	writeln('Resta: ',resta);

	readln;
end.

