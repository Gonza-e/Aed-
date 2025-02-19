program a; 

uses 
	crt;
var 
	N,i,j: integer;
begin
	writeln('Ingrese un numero mayor a 3: ');
	readln(N); 

	clrscr;
	for i:= 0 to N-1 do 
	begin 
		for j:= 0 to N-1 do 
		begin 
			if (((i=j) and (j>n div 2)) or (j=n div 2)) or (((i+j=n-1) and (j > n div 2)) or (j=n div 2)) then //Si queremos hacer el logo de la utn solo borramos (and j > n div 2)
			begin 
				write('*'); 
			end 
			else 
			begin 
				write(' '); 
			end;
		end; 
		writeln();
	end; 

	readln; 
end. 