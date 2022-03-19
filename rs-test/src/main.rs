struct Student{
	name         :String,
	cLanguage    :f32,
	javaLanguage :f32,
}

fn main(){
	let students = vec![
		Student{name: "나".to_string(), cLanguage: 100., javaLanguage: 100.},
		Student{name: "철수".to_string(), cLanguage: 90., javaLanguage: 80.},
		Student{name: "영희".to_string(), cLanguage: 90., javaLanguage: 95.},
    ];
    let mut data = students.into_iter()
        .map(|student|(student.name, (student.cLanguage + student.javaLanguage) / 2.))
        .filter(|(name, mean)| *mean >= 90.)
        .collect::<Vec<_>>();
    data.sort_by(|(a_name, a_mean), (b_name, b_mean)|a_mean.partial_cmp(b_mean).unwrap());
    data.iter().for_each(|(name, mean)|{
        println!("{} : {}", name, mean);
    });
}