Program Recursividad; 

 
function verificar(numero: LongInt; cont_ceros: integer): boolean; 
begin
	if numero = 0 then 
	begin
		if ((cont_ceros mod 3) = 0) then 
		begin
			verificar:= True;
		end 
		else 
		begin 
			verificar:= False; 
		end;
	end
	else 
	begin 
		if ((numero mod 10) = 0) then 
		begin 
			verificar:= verificar(numero div 10, cont_ceros + 1);
		end
		else
		begin
			verificar:= verificar(numero div 10, cont_ceros);
		end;
	end;
end;

var 
	binario: LongInt;
	ceros: integer;
begin
	binario:= 1001010;
	ceros:= 0;

	if verificar(binario,ceros) then 
	begin
		writeln('El numero tiene una cantidad multiplo de 3 de ceros');
	end
	else
	begin 
		writeln('El numero no cumple con la condicion');
	end;

	readln; 
end.