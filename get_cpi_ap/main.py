from tabula import read_pdf
from tabulate import tabulate
import pandas as pd
import os
import json

#
#   grade_structure {
#       "sems" : [
#           "courses": [
#               "course": Course_ID, 
#               "credits": credits, 
#               "grade": grade_obtained, 
#               "credits_received": credits
#           ]
#           "spi": val
#           "sem_num": n
#           "credits_done": n
#       ]
#       "status": Normal
#   }

def return_pandas_df(tabl, name_file):
    # print("one file")
    
    location = "C:/Users/fahee/programming/webdev/ap_cpi_calc/get_cpi_ap"
    tabl.to_csv(name_file)
    pd_df = pd.read_csv(name_file)
    path = os.path.join(location, name_file)
    os.remove(os.path.join(name_file))
    return pd_df

def isinteg(num):
    try:
        int(num)
        return True
    except ValueError:
        return False

def isfloat(num):
    try:
        float(num)
        return True
    except ValueError:
        return False

def acad_status_calc(prev_sem_tot_credits, num_sems, total_credits, prev_sem_status):
    status = "Normal"
    if prev_sem_tot_credits >= 30 and (total_credits >= num_sems*(num_sems+24) and total_credits < 36*num_sems):
        status = "Warning"
    if prev_sem_tot_credits < 30 and (total_credits >= 36*num_sems):
        status = "Warning"

    if prev_sem_tot_credits >= 30 and (total_credits < (24+num_sems)*num_sems):
        status = "Academic Probation"
    if prev_sem_tot_credits < 30 and (total_credits >= num_sems*(num_sems+24) and total_credits < 36*num_sems):
        status = "Academic Probation"
    if prev_sem_tot_credits < 30 and (total_credits < (24+num_sems)*num_sems) and prev_sem_status != "Academic Probation":
        status = "Academic Probation"
    if prev_sem_status == "Academic Probation" and prev_sem_tot_credits < 30 and total_credits < (24 + num_sems)*num_sems:
        status = "Programme Termination"
    return status


def get_cpi_ap(file):
    #reads table from pdf file
    df = read_pdf(file,pages="all") #address of pdf file

    # tabula.convert_into("transcript2.pdf", "output.csv", output_format="csv", pages='all')

    grade_structure = {}
    grade_structure["sems"] = []

    total_credits = 0
    final_cpi = 0
    temp_numer = 0
    temp_denom = 0
    count = 0
    prev_sem_tot_credits = 0
    num_sems = 0
    prev_sem_status = "Normal"
    one_sem = {}
    one_sem["courses"] = []
    # courses_one_sem = []
    for tabl in df:
        if len(tabl.columns) > 5:
            name = "table_" + str(count) + ".csv"
            pd_df = return_pandas_df(tabl, name)
            count = count + 1
            for row in pd_df.itertuples(index=True, name='Pandas'):
                row = row[::-1]
                # print(row)
                credits = 0
                flag = 0
                scored = 0
                grade = ''
                for val in row:
                    val = str(val)
                    if val != 'nan':
                        if 'CPI/SPI' in val:
                            # print("numer: ", temp_numer, " den: ", temp_denom)
                            if temp_denom != 0:
                                # print("SPI is: ",temp_numer/temp_denom*10)
                                prev_sem_status = acad_status_calc(prev_sem_tot_credits, num_sems, total_credits - temp_denom, prev_sem_status)
                                num_sems = num_sems + 1
                                sem_number = "sem " + str(num_sems)
                                # one_sem["courses"] = courses_one_sem
                                # courses_one_sem = {}
                                one_sem["spi"] = temp_numer/temp_denom*10
                                one_sem["sem_num"] = num_sems
                                one_sem["credits_done"] = temp_denom
                                grade_structure["sems"].append(one_sem)
                                one_sem = {}
                                one_sem["courses"] = []
                            prev_sem_tot_credits = temp_denom
                            temp_denom = 0
                            temp_numer = 0
                            break
                        if flag == 1:
                            if len(val) >= 5 and len(val) <= 8:
                                if any(chr.isdigit() for chr in val):
                                    one_sem["courses"].append({"course": val, "credits": credits, "grade": grade, "credits_received": scored})
                                    break
                            continue
                        if credits != 0:
                            flag = 1
                            grade = val
                            match val:
                                case 'A*':
                                    scored = credits
                                case 'A':
                                    scored = credits
                                case 'B':
                                    scored = credits*0.8
                                case 'C':
                                    scored = credits*0.6
                                case 'D':
                                    scored = credits*0.4
                                case 'E':
                                    scored = credits*0.2
                                case 'F':
                                    scored = 0
                                case 'S':
                                    total_credits = total_credits + credits
                                    break
                                case default:
                                    break
                            total_credits = total_credits + credits
                            temp_denom = temp_denom + credits
                            temp_numer = temp_numer + scored
                            continue
                        if isinteg(val):
                            credits = int(val)
                        if isfloat(val):
                            tem = float(val)
                            credits = int(tem)
                        # print("denom: ", temp_denom)

    if temp_denom != 0:
        # print("SPI is: ",temp_numer/temp_denom*10)
        num_sems = num_sems + 1
        prev_sem_tot_credits = temp_denom
        sem_number = "sem " + str(num_sems)
        one_sem["spi"] = temp_numer/temp_denom*10
        grade_structure[sem_number] = one_sem
    status = acad_status_calc(prev_sem_tot_credits, num_sems, total_credits, prev_sem_status)
    grade_structure["status"] = status
    return grade_structure
