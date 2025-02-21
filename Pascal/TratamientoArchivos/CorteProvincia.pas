Program CorteProvincia; 

type 
  PROVINCIAS = record
    provincia: String[30];
    nombre_jefe: String[30];
    oficinas: 1..6;
    cant_empleados: Integer;
  end;
var 
 arch: File of PROVINCIAS;
 reg: PROVINCIAS;
 resg_provincia, resg_jefe: String[30];
 resg_oficina: 1..6;
 cant_oficina, cant_jefe, cant_provincia, total: Integer;

procedure corte_oficina();
begin
   WriteLn('En la oficina ',resg_oficina,' hay ',cant_oficina,' empleados');
   cant_jefe:= cant_jefe + cant_oficina;
   cant_oficina:= 0; 
   resg_oficina:= reg.oficinas;
   resg_jefe:= reg.nombre_jefe;
   resg_provincia:= reg.provincia;
end;

procedure corte_jefe(); 
begin
   corte_oficina;
   WriteLn('El jefe ',resg_jefe,' tiene ',cant_jefe,' empleados...');
   cant_provincia:= cant_provincia + cant_jefe;
   cant_jefe:= 0;
   resg_jefe:= reg.nombre_jefe;
   resg_provincia:= reg.provincia;
end;

procedure corte_provincia();
begin
   corte_jefe;
   WriteLn('La provincia ',resg_provincia,' tiene ',cant_provincia,' empleados___');
   total:= total + cant_provincia;
   cant_provincia:= 0;
   resg_provincia:= reg.provincia;
end;

begin
  Assign(arch,'provincias.dat');
  Reset(arch);
  if not Eof(arch) then
  begin 
    Read(arch,reg);
    resg_provincia:= reg.provincia;
    resg_jefe:= reg.nombre_jefe;
    resg_oficina:= reg.oficinas;
  end;
  cant_oficina:= 0; cant_jefe:= 0; cant_oficina:= 0; 
  while not Eof(arch) do
  begin
   cant_oficina:= cant_oficina + reg.cant_empleados;
    if resg_provincia <> reg.provincia then
     begin
       corte_provincia;
     end
     else if resg_jefe <> reg.nombre_jefe then
     begin
       corte_jefe;
     end
     else if resg_oficina <> reg.oficinas then
     begin
       corte_oficina;
     end;
     Read(arch,reg)
  end;
 corte_provincia;
 WriteLn('La cantidad total de empleados es de: ',total);
 Close(arch);
 ReadLn;
end.
