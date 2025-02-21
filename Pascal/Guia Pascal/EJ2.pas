Program Ejercicio2;

{2) Idem al ejercicio anterior, pero ahora tanto x como y deben ser ingresados por el usuario}

var
	x, y, resta, resto, division, suma, mul: integer; 
begin
	writeln('Ingrese el primer valor'); //Uso el readln para entradas de teclado ya que el read solo lee el valor 
	readln(x);
	writeln('Ingrese el segundo valor');
	readln(y);

	resta:= x - y; 
	resto:= x mod y; 
	division:= x div y; 
	suma:= x + y; 
	mul:= x * y;

	writeln('Resta: ',resta);
	writeln('Division: ',division);
	writeln('Suma: ',suma);
	writeln('Resto: ',resto);
	writeln('Multiplicacion: ',mul);

	readln;
end.