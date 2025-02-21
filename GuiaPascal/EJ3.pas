Program Ejercicio1;

{3) Idem al anterior pero ahora el usuario es el que ingresa que operación quiere hacer}

var 
	resultado, operacion, x, y: integer;

	procedure mostrar_operacion(); 
	begin
		case operacion of 
			1: resultado:= x + y;
			2: resultado:= x div y;
			3: resultado:= x * y;
			4: resultado:= x mod y;
			5: resultado:= x - y;
		end; 

		writeln('El resultado es: ',resultado);
	end;
begin
	writeln('Ingrese el primer valor:');
	readln(x);
	writeln('Ingrese el segundo valor:');
	readln(y);
	writeln('Ingrese operacion: 1 = suma, 2 = division entera, 3 = multiplicacion, 4 = resto, 5 = resta');
	readln(operacion);
	mostrar_operacion(); 

	readln;
end.
