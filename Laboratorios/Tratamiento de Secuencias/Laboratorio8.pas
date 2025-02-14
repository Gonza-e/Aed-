Program Laboratorio8; 

{El Registro automotor posee datos de autos en el archivo REGISTRO_AUTO.txt.
El mismo contiene las patentes (longitud indeterminada, finaliza con '&')
y las fechas de inscripción de cada uno de los automóviles (8 caracteres que representan: ddmmaaaa).

Ejemplo: LZR313&29082021AB446HF&06021990...

PREGUNTA: Indique el porcentaje de autos inscriptos antes del año 2000.}


uses 
 math, 
 sysutils; 
var 
 sec: TextFile;
 v: Char; 
 autos_total, autos_cumple, i: Integer;
 anio: String;
 anio_int: Integer;
 porc: Real;

 function conv_entero(c:Char):Integer;
 begin
   case c of 
     '1': conv_entero:= 1;
     '2': conv_entero:= 2;
     '3': conv_entero:= 3;
     '4': conv_entero:= 4;
     '5': conv_entero:= 5;
     '6': conv_entero:= 6;
     '7': conv_entero:= 7;
     '8': conv_entero:= 8;
     '9': conv_entero:= 9;
     '0': conv_entero:= 0;
    end;
 end;
begin 
  Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\registro_auto.txt');
  Reset(sec); 
  Read(sec,v); 

  autos_total:= 0; autos_cumple:= 0; porc:= 0; anio_int:= 0;
  anio:= ' ';
  while not Eof(sec) do
  begin

    while (not Eof(sec)) and (v <> '&') do
    begin
      Read(sec,v);
    end;
    Read(sec,v); 
   
    if not Eof(sec) then
    begin

        Read(sec,v);
        Read(sec,v);
        Read(sec,v);
        Read(sec,v);
    
        {for i:= 4 downto 1 do 
        begin
        if (anio = ' ') then
        begin
            anio:= v;
        end else 
        begin
            anio:= anio + v;  
        end;
        Read(sec,v);
        //ensec:= True;
        end;}

        anio_int:= anio_int + conv_entero(v) * 1000;
        Read(sec,v);
        anio_int:= anio_int + conv_entero(v) * 100;
        Read(sec,v);
        anio_int:= anio_int + conv_entero(v) * 10;
        Read(sec,v);
        anio_int:= anio_int + conv_entero(v);

        //anio_int:= strtoint(anio);
        WriteLn(anio_int);  
        autos_total+=1;

        if anio_int < 2000 then 
        begin
        autos_cumple+=1;
        end;
        anio_int:= 0;
        //anio:= ' ';

        {WriteLn(autos_total);
        WriteLn(autos_cumple);}
    end;

  end;
  Close(sec);
  porc:= (autos_cumple/autos_total) * 100;
  WriteLn('El porcentaje de autos que cumplen con la condicion es de: ',porc:0:2,'%');
end. 