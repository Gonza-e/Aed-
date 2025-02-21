Program CorteSolicitudes; 

type 
  SOLICITUDES = record 
    usuario: String[20];
    tipo: String[2];
    cant_solicitudes: Integer;
    restringido: String[2];
  end; 

var 
  arch, arch_sal: File of SOLICITUDES;
  reg, reg_sal: SOLICITUDES;
  cant_tipo, cant_usu, total: integer; 
  resg_tipo: String[2];
  resg_usuario: String[20];

procedure nuevo_archivo();
begin
  reg_sal.usuario:= reg.usuario;
  reg_sal.cant_solicitudes:= reg.cant_solicitudes;
  Write(arch_sal,reg_sal);
end;

procedure corte_tipo(); 
begin
  WriteLn('La cantidad de solicitudes de tipo ',resg_tipo,' es de ',cant_tipo);
  cant_usu:= cant_usu + cant_tipo;
  cant_tipo:= 0;
  if resg_tipo = 'S' then
  begin
    nuevo_archivo;
  end;
  resg_tipo:= reg.tipo;
  resg_usuario:= reg.usuario;
end;

procedure corte_usuario();
begin
  corte_tipo;
  WriteLn('La cantidad de solicitudes del usuario ',resg_usuario,' es de ',cant_usu);
  total:= total + cant_usu;
  cant_usu:= 0;
  resg_usuario:= reg.usuario;
end;


begin
  Assign(arch,'solicitudes.dat');
  Reset(arch);
  Assign(arch_sal,'archivo salida');
  Rewrite(arch_sal);
  cant_tipo:= 0; cant_usu:= 0; 
  if not Eof(arch) then
  begin
    Read(arch,reg);
    resg_tipo:= reg.tipo;
    resg_usuario:= reg.usuario;
  end;
  while not Eof(arch) do
  begin
    cant_tipo:= cant_tipo + reg.cant_solicitudes;
    if resg_usuario <> reg.usuario then
    begin
      corte_usuario;
    end
    else if resg_tipo <> reg.tipo then 
    begin
      corte_tipo;
    end;
    Read(arch,reg);
  end;
  corte_usuario;
  WriteLn('La cantidad total de solicitudes es de: ',total);
  Close(arch);
  Close(arch_sal);
  ReadLn;
end.