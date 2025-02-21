Program Provincias; 

{contar cuantos jefes tienen mas de 100 empleados y cuales son de la planta 4}

uses 
 csvdocument,
 sysutils; 
type 
  PROVINCIA = record
   provincias: String[30];
   nombre_jefe: String[30];
   oficinas: 1..6;
   cant_empleados: Integer;
  end;
var 
 arch: TCSVDocument;
 reg: PROVINCIA;
 cant_100, cant_planta4, i: Integer;
 procedure INICIALIZAR();
    begin
      cant_100:=0;
      cant_planta4:=0;
      arch:= TCSVDocument.create();//aca se crea el archivo
    end;
begin
  INICIALIZAR();
  arch.delimiter:= ';';
  arch.LoadFromFile('C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\TratamientoArchivos\Archivos\provincias.csv');
  for i:= 1 to (arch.rowcount - 1) do
    begin
        reg.provincias:=arch.cells[0,i];
        reg.nombre_jefe:=arch.cells[1,i];
        reg.oficinas:=StrToInt(arch.cells[2,i]);
        reg.cant_empleados:=StrToInt(arch.cells[3,i]);

        if (reg.cant_empleados >= 100) then     
        begin
            cant_100+=1;
            WriteLn('Jefe con mas de 100: ',reg.nombre_jefe);
        end; 
        if (reg.oficinas = 4) then 
        begin
          cant_planta4+=1;
          WriteLn('Jefe de oficina 4: ',reg.nombre_jefe);
        end;   
    end;
    WriteLn('la cantidad de jefes que tiene mas de 100 empleado es: ',cant_100);
    WriteLn('la cantidad de jefes de la planta 4 es de: ',cant_planta4);
    arch.free;
    ReadLn;
end.

