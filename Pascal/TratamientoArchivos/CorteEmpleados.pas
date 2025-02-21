Program CorteEmpleados; 

type
  EMPLEADOS = record
    planta: String[15];
    categoria: String[15]; 
    nombre: string[50];
  end;
var 
 arch: file of EMPLEADOS;
 reg: EMPLEADOS;
 cant_cat, cant_planta, total: Integer;
 resg_planta, resg_cat: String[15];

procedure corte_categoria();
begin
  WriteLn('|La cantidad de empleados de la categoria: ',resg_cat,'| |es de: ',cant_cat,'|');
  cant_planta:= cant_planta + cant_cat;
  cant_cat:= 0;
  resg_cat:= reg.categoria
end;

procedure corte_planta();
begin
  corte_categoria;
  WriteLn('|La cantidad de empleados en la planta: ',resg_planta,'| |es de: ',cant_planta,'|');
  total:= total + cant_planta;
  cant_planta:= 0;
  resg_cat:= reg.categoria;
  resg_planta:= reg.planta;
end;

begin
  Assign(arch, 'empleados1.dat');
  Reset(arch);
  cant_cat:= 0;
  cant_planta:= 0;
  if not Eof(arch) then
  begin
    Read(arch, reg);
    resg_cat := reg.categoria;
    resg_planta := reg.planta;
  end;
  while not Eof(arch) do
  begin
   WriteLn('Nombre del empleado: ',reg.nombre);
    if resg_planta <> reg.planta then
    begin
      corte_planta;
    end
    else
    begin
      if resg_cat <> reg.categoria then
        corte_categoria;
    end;
    cant_cat:= cant_cat + 1;
    Read(arch,reg)
  end;

  corte_planta;

  WriteLn('La cantidad total de empleados es de: ', total);
  Close(arch);
  ReadLn;
end.
