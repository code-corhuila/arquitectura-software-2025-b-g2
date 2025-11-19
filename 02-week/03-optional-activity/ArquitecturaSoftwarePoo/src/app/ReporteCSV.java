package app;

import java.util.List;

public class ReporteCSV implements Reporte {
    @Override
    public String generar(List<Empleado> empleados) {
        StringBuilder sb = new StringBuilder("id,nombre,salario\n");
        for (Empleado e : empleados) {
            sb.append(e.getId()).append(",")
                    .append(e.getNombre()).append(",")
                    .append(e.calcularSalario()).append("\n");
        }
        return sb.toString();
    }
}