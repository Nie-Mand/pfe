# PFE, Please

Send an email requesting a PFE Internship to all the companies in the list.
It is based on the Courier API, so it is required to have an account with a valid provider and valid template and brand.

## Usage

1. Install the binary

```bash
```


2. Create `config.txt` file, it should contain the following information:
    - `COURIER_TOKEN`: your [Courier](https://courier.com) token
    - `TEMPLATE_ID`: the template id of the email you want to send
    - `BRAND_ID`: the brand id of the email you want to send

```env
COURIER_TOKEN=your_token
TEMPLATE_ID=your_template_id
BRAND_ID=your_brand_id
```

3. Create `companies.csv` file, it should contain the following information:
    - `company`: the name of the company
    - `email`: the email of the company
    - `position`: the position you are applying for

```csv
company,position,email
Google,Software Engineering Intern,hr@google.com
Tiktok,Platform Engineering Intern,hr@tiktok.com
```

4. Run the script

```bash
./pfe send -f companies.csv
```