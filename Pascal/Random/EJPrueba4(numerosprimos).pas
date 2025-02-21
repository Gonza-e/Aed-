Program a; 

uses 
	crt, math, sysutils; 
var 
	num, i: integer;

	function EsPrimo(x: integer): boolean;
	var
	  j: integer;
	begin
	  if (x < 2) then
	    EsPrimo := false
	  else
	  begin
	    EsPrimo := true;
	    for j := 2 to trunc(sqrt(x)) do
	       	if (x mod j = 0) then
	     	begin
        		EsPrimo := false;
        		break;
      		end;
  		end;
	end;
begin
	writeln('Ingrese numero: ');
	readln(num); 

	for i:= 1 to (num-1) do  
	begin 
		if EsPrimo(i) then  
		begin 
			writeln(i);
		end; 
	end;

	readln;  
end. 