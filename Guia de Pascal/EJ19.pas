Program EJ19; 

{5) idem al ejercicio anterior pero ahora la multiplicación de cada dígito de la contraseña
debe ser por la posición en la que se encuentra en el número. El primer dígito se multiplica
por 1, el segundo por 2, el tercero por 3 y así.
Ej: contraseña igual a 5362, va a ser 5 x 1 + 3 x 2 + 6 x 3 + 2 x 4
La contraseña “válida”, al igual que el ejercicio 4, es la que el resultado es mayor a los 2
números menos significativos del usuario.}

uses 
	sysutils; 
var 
	USUARIO_clave, USUARIO_num: LongInt; 
	CLAVE_string: string;
	CLAVE_int, i, ULTDOS_num: integer;

	procedure validar();
	begin
		for i:= length(CLAVE_string) downto 1 do  
		begin
			CLAVE_int:= CLAVE_int + ((USUARIO_clave mod 10)*i);
			USUARIO_clave:= USUARIO_clave div 10;
		end; 

		if CLAVE_int > ULTDOS_num then 
		begin
			writeln('La clave es valida'); 
		end
		else 
		begin
			writeln('La clave es invalida'); 
		end;
	end;

begin
	writeln('Ingrese el numero de usuario (5 numeros)');
	readln(USUARIO_num);

	writeln('Ingrese la clave de usuario');
	readln(USUARIO_clave); 

	CLAVE_string:= IntToStr(USUARIO_clave);
	ULTDOS_num:= USUARIO_num mod 100;

	validar();
	readln;
end.  
