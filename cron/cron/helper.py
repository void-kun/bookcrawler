from datetime import datetime


def convert_date(date: str) -> str:
    str_date = datetime.strptime(date, '%d-%m-%Y')
    return str_date.strftime('%Y-%m-%d')


def fstr(template: str, **kwargs) -> str:
    return eval(f"f'{template}'", kwargs)
