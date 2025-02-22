Program Recursividad1; 

function verificar(numero, cont_patron: LongInt): boolean;
begin
	if numero = 0 then 
	begin
		if cont_patron > 0 then 
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
		if ((numero mod 100) = 10) then 
		begin
			verificar:= verificar(numero div 100, cont_patron + 1);
		end 
		else 
		begin
			verificar:= verificar(numero div 100, cont_patron);
		end;
	end;
end;

begin

	if verificar(11000110,0) then 
	begin
		writeln('El numero cumple con el patron');
	end
	else 
	begin
		writeln('El numero no cumple con el patron');
	end;

	readln;
end.