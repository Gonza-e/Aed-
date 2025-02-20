program a; 

uses 
	crt;
var 
	i,num: integer;
	fact: longint;

	procedure factorial();
	begin
		fact:= 1;
		for i:= 1 to num do 
		begin
			fact:= fact*i;
		end;  
	end;   
begin
	writeln('Ingrese el numero para calcular factorial');
	readln(num);  

	factorial();
	clrscr;

	writeln('El factorial de ',num,' es ',fact);
	readln;
end. 