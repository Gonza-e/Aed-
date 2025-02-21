Program CorteProvincias2; 

{Mostrar por pantalla cantidad de empleados por provincia y por jefe}

uses 
 csvdocument,
 sysutils;
type
  PROVINCIAS = record
    provincia: String[30];
    nombre_jefe: String[30];
    oficinas: 1..6;
    cant_empleados: Integer;
  end;
var 
 arch: TCSVDocument; 
 reg: PROVINCIAS; 
 resg_provincia, resg_jefe: String[30]; 
 empleados_jefe, empleados_provincia, total, i: Integer; 
 
 procedure INICIALIZAR();
 begin
   empleados_jefe:= 0; empleados_provincia:= 0; total:= 0;
   arch:= TCSVDocument.create();
 end;

 procedure corte_jefe();
 begin
   WriteLn('El jefe ',resg_jefe,' tiene ',empleados_jefe,' empleados');
   empleados_provincia:= empleados_provincia + empleados_jefe;
   empleados_jefe:= 0; 
   resg_jefe:= arch.cells[1,i];
 end;

 procedure corte_provincia();
 begin
   corte_jefe();
   WriteLn('La provincia ',resg_provincia,' tiene ',empleados_provincia,' empleados');
   total:= total + empleados_provincia;
   empleados_provincia:= 0; 
   resg_provincia:= arch.cells[0,i];
 end;
begin
  INICIALIZAR();
  arch.delimiter:= ';';
  arch.LoadFromFile('C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\TratamientoArchivos\Archivos\provincias2.csv');
  resg_provincia:= arch.cells[0,1]; //arch.cells[j,i] donde 'j' son las columnas e 'i' son las filas
  resg_jefe:= arch.cells[1,1];
  
  for i:= 1 to (arch.rowcount - 1) do 
  begin
    reg.provincia:= arch.cells[0,i];
    reg.nombre_jefe:= arch.cells[1,i];
    reg.oficinas:= StrToInt(arch.cells[2,i]);
    reg.cant_empleados:= StrToInt(arch.cells[3,i]);
    
    empleados_jefe:= empleados_jefe + reg.cant_empleados;
    if resg_provincia <> reg.provincia then
    begin
      corte_provincia();
    end
    else if resg_jefe <> reg.nombre_jefe then 
    begin
      corte_jefe();
    end;
  end;
  corte_provincia();
  arch.free;
  WriteLn('La cantidad de empleados en total es de: ',total);
  ReadLn;
end.
