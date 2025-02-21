Program EJ11; 

{11) Se pide crear un pequeño login para usuarios, el usuario tiene que pasar su nombre y
su contraseña, el sistema tiene que comparar contra unos valores previamente cargados y
verificar si son o no iguales.}

var 
	usu, pass: string;
const
	usuario = 'Usuario';
	contrasena = 'Password'; 
begin
	writeln('Ingrese nombre de usuario: ');
	readln(usu); 
	writeln('Ingrese contrasena: ');
	readln(pass); 

	if (contrasena = pass) and (usuario = usu) then 
	begin  
		writeln('Usuario y contrasena correcta');
	end
	else 
		writeln('Usuario y contrasena incorrectas');

	readln;
end. 