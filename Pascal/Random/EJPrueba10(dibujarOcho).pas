program a; 

uses
	crt, sysutils; 
var 
	i,j: integer;
begin 
	for i:= 1 to 5 do 
	begin 
		for j:= 1 to 4 do 
		begin 
			if (i=1) or (i=3) or (i=5) or ((j=4) and (i=4)) or (((j=4) or (j=1)) and (i=2)) then 
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