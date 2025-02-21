Program EJ16;

{3) Escribir un algoritmo donde se ingresan 3 números distintos y se deben ordenar de
menor a mayor, e imprimir por pantalla.}

var 
	A: array[1..3] of integer;
	min, v_min, i, j: integer; 
begin
	writeln('Ingrese numero 1');
	readln(A[1]);
	writeln('Ingrese numero 2');
	readln(A[2]);
	writeln('Ingrese numero 3');
	readln(A[3]);	

	for i:= 1 to 2 do 
	begin
		min:= A[i];
		for j:= i+1 to 3 do 
		begin
			if A[j] < min then 
			begin
				v_min:= j;
				min:= A[j];
			end; 
		end;
		A[v_min]:= A[i];
		A[i]:= min; 
	end;
	
	for i:= 1 to 3 do 
	begin
		writeln('Numero ',i,':',A[i]); 
	end;	

 	readln;
end.