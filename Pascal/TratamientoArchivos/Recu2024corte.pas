Program Recu2024corte; 

type
  ATLETAS = record 
    delegacion: String[20];
    deporte: String[3];
    nrocredencial: String[20];
    apeynom: String[20];
    participaciones: Integer;
  end;

  RESUMEN = record
    delegacion: String[20];
    con_experiencia: Integer;
    sin_experiencia: Integer;
  end;

var 
  arch_sal: File of RESUMEN;
  reg_sal: RESUMEN;
  arch: File of ATLETAS;
  reg: ATLETAS;
  conexp, sinexp, conexp_depor, sinexp_depor, cant_deporte, cant_deleg, total: Integer;
  resg_delegacion, resg_deporte: String[20];

procedure corte_deporte();
begin
  WriteLn('|Delegacion: ', resg_delegacion, '| |Deporte: ', resg_deporte, '| |Cantidad: ', cant_deporte,'|');
  cant_deleg := cant_deleg + cant_deporte;
  cant_deporte := 0; 
  resg_deporte := reg.deporte;
end;

procedure corte_delegacion();
begin
  corte_deporte; 
  WriteLn('|Delegacion: ', resg_delegacion, '| |Cantidad: ', cant_deleg,'|');
  reg_sal.delegacion:= resg_delegacion;
  reg_sal.con_experiencia:= conexp_depor;
  reg_sal.sin_experiencia:= sinexp_depor;
  Write(arch_sal,reg_sal);
  conexp_depor:= 0;
  sinexp_depor:= 0;
  total := total + cant_deleg;
  cant_deleg := 0;
  resg_delegacion := reg.delegacion; 
end;

begin
  Assign(arch,'atletas.dat');
  Reset(arch); 
  Assign(arch_sal,'archivoResumen');
  Rewrite(arch_sal);
  cant_deporte := 0; cant_deleg := 0; total := 0; conexp := 0; sinexp := 0; conexp_depor:= 0; sinexp_depor:= 0;

  if not Eof(arch) then 
  begin
    Read(arch, reg);
    resg_delegacion := reg.delegacion;
    resg_deporte := reg.deporte;
  end;

  while not Eof(arch) do
  begin

    if resg_delegacion <> reg.delegacion then
    begin
      corte_delegacion;
    end
    else if resg_deporte <> reg.deporte then
    begin
      corte_deporte;
    end; 

    if reg.participaciones <= 1 then
    begin
      sinexp := sinexp + 1;
      sinexp_depor:= sinexp_depor + 1;
    end
    else
    begin
      conexp := conexp + 1;  
      conexp_depor:= conexp_depor + 1;
    end;

    cant_deporte := cant_deporte + 1;

    Read(arch, reg);  
  end;

  corte_delegacion;  

  WriteLn('La cantidad de atletas con experiencia es de: ', conexp);
  WriteLn('La cantidad de atletas sin experiencia es de: ', sinexp);
  WriteLn('Cantidad total de atletas: ', total);

  Close(arch);
  Close(arch_sal);
  ReadLn;
end.
