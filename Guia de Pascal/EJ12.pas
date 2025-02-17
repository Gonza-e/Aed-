Program EJ12; 

{12) Dado el ejercicio anterior, realizar el mismo algoritmo, solo que ahora no se muestre el
valor de la contraseña cuandos se escribe por el teclado.}

uses 
	crt; 
var 
	usuario, pass: string;
	letra: char;
begin 
	writeln('Ingrese el nombre de usuario: ');
	readln(usuario);
	writeln('Ingrese la password'); 
	pass:= ' '; 

	while (letra <> #13) do 
	begin
		letra:= readkey; 

		if (letra <> #13) then
		begin
			if pass <> ' ' then
			begin
				pass:= pass + letra;
			end 
			else 
			begin 
				pass:= letra; 
			end;
			write(chr(31));  
		end;  
	end; 

	writeln();
	writeln('El nombre de usuario es: ',usuario);
	writeln('La password es: ',pass);
	readln;
end. 